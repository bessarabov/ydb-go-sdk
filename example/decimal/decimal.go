package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"math/big"
	"path"
	"text/template"

	"github.com/yandex-cloud/ydb-go-sdk"
	"github.com/yandex-cloud/ydb-go-sdk/decimal"
	"github.com/yandex-cloud/ydb-go-sdk/example/internal/cli"
	"github.com/yandex-cloud/ydb-go-sdk/table"
)

type templateConfig struct {
	TablePathPrefix string
}

var writeQuery = template.Must(template.New("fill database").Parse(`
PRAGMA TablePathPrefix("{{ .TablePathPrefix }}");

DECLARE $decimals AS List<Struct<
	id: Uint32,
	value: Decimal(22,9)>>;

REPLACE INTO decimals
SELECT
	id,
	value
FROM AS_TABLE($decimals);
`))

var readQuery = template.Must(template.New("fill database").Parse(`
PRAGMA TablePathPrefix("{{ .TablePathPrefix }}");
SELECT value FROM decimals;
`))

type Command struct {
	config func(cli.Parameters) *ydb.DriverConfig
}

func (cmd *Command) ExportFlags(ctx context.Context, flag *flag.FlagSet) {
	cmd.config = cli.ExportDriverConfig(ctx, flag)
}

func (cmd *Command) Run(ctx context.Context, params cli.Parameters) error {
	driver, err := (&ydb.Dialer{
		DriverConfig: cmd.config(params),
	}).Dial(ctx, params.Endpoint)
	if err != nil {
		return err
	}

	c := table.Client{Driver: driver}
	session, err := c.CreateSession(ctx)
	if err != nil {
		return err
	}
	defer session.Close(context.Background())

	var (
		tablePathPrefix = path.Join(params.Database, params.Path)
		tablePath       = path.Join(tablePathPrefix, "decimals")
	)
	err = session.CreateTable(ctx, tablePath,
		table.WithColumn("id", ydb.Optional(ydb.TypeUint32)),
		table.WithColumn("value", ydb.Optional(ydb.DefaultDecimal)),
		table.WithPrimaryKeyColumn("id"),
	)
	if err != nil {
		return err
	}

	write, err := session.Prepare(ctx, render(writeQuery, templateConfig{
		TablePathPrefix: tablePathPrefix,
	}))
	if err != nil {
		return err
	}
	read, err := session.Prepare(ctx, render(readQuery, templateConfig{
		TablePathPrefix: tablePathPrefix,
	}))
	if err != nil {
		return err
	}

	txc := table.TxControl(
		table.BeginTx(
			table.WithSerializableReadWrite(),
		),
		table.CommitTx(),
	)

	x := big.NewInt(42 * 1000000000)
	x.Mul(x, big.NewInt(2))

	_, _, err = write.Execute(ctx, txc, table.NewQueryParameters(
		table.ValueParam("$decimals",
			ydb.ListValue(
				ydb.StructValue(
					ydb.StructFieldValue("id", ydb.Uint32Value(42)),
					ydb.StructFieldValue("value", ydb.DecimalValue(
						ydb.DefaultDecimal,
						decimal.Int128(x, 22, 9),
					)),
				),
			),
		),
	))
	if err != nil {
		return err
	}

	_, res, err := read.Execute(ctx, txc, nil)
	if err != nil {
		return err
	}
	for res.NextSet() {
		for res.NextRow() {
			res.NextItem()
			p := res.ODecimal(ydb.DefaultDecimal)

			x := decimal.FromInt128(p, 22, 9)
			fmt.Println(decimal.Format(x, 22, 9))
		}
	}
	return res.Err()
}

func render(t *template.Template, data interface{}) string {
	var buf bytes.Buffer
	err := t.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

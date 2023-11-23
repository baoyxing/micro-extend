package entutils

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"log"
	"time"
)

const (
	mysqlEventKey = "mysql"
)

var (
	mysqlQueryTextKey = attribute.Key("otel.mysql.query")
	mysqlArgsTextKey  = attribute.Key("otel.mysql.args")
	mysqlDurationKey  = attribute.Key("otel.mysql.duration")
)

type EntMysqlOption interface {
	apply(driver *EntMysqlDriver)
}

type entMysqlOption func(driver *EntMysqlDriver)

func (fn entMysqlOption) apply(driver *EntMysqlDriver) {
	fn(driver)
}

type EntMysqlDriver struct {
	dialect.Driver                               // underlying driver.
	log            func(context.Context, ...any) // log function. defaults to log.Println.
	isTrace        bool                          // whether to trace queries.
}

func NewMysqlDriver(driver dialect.Driver, options ...EntMysqlOption) *EntMysqlDriver {
	entDriver := DefaultMysqlDriver(driver)
	for _, option := range options {
		option.apply(entDriver)
	}
	return entDriver
}

func DefaultMysqlDriver(driver dialect.Driver) *EntMysqlDriver {
	return &EntMysqlDriver{
		Driver: driver,
		log: func(ctx context.Context, info ...any) {
			log.Println(info...)
		},
		isTrace: false,
	}
}
func WithLog(log func(context.Context, ...any)) EntMysqlOption {
	return entMysqlOption(func(driver *EntMysqlDriver) {
		if log != nil {
			driver.log = log
		}

	})
}

func WithTrace(isTrace bool) EntMysqlOption {
	return entMysqlOption(func(driver *EntMysqlDriver) {
		driver.isTrace = isTrace
	})
}

// Exec logs its params and calls the underlying driver Exec method.
func (d *EntMysqlDriver) Exec(ctx context.Context, query string, args, v any) error {
	start := time.Now()
	err := d.Driver.Exec(ctx, query, args, v)
	duration := time.Since(start)
	d.startTrace(ctx, query, duration, args)
	d.log(ctx, fmt.Sprintf("driver.Exec: query=%v args=%v time=%v", query, args, duration))
	return err
}

// ExecContext logs its params and calls the underlying driver ExecContext method if it is supported.
func (d *EntMysqlDriver) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	drv, ok := d.Driver.(interface {
		ExecContext(context.Context, string, ...any) (sql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	start := time.Now()
	result, err := drv.ExecContext(ctx, query, args...)
	duration := time.Since(start)
	d.startTrace(ctx, query, duration, args)
	d.log(ctx, fmt.Sprintf("driver.ExecContext: query=%v args=%v time=%v", query, args, duration))
	return result, err
}

// Query logs its params and calls the underlying driver Query method.
func (d *EntMysqlDriver) Query(ctx context.Context, query string, args, v any) error {
	start := time.Now()
	err := d.Driver.Query(ctx, query, args, v)
	duration := time.Since(start)
	d.startTrace(ctx, query, duration, args)
	d.log(ctx, fmt.Sprintf("driver.Query: query=%v args=%v time=%v", query, args, duration))
	return err
}

// QueryContext logs its params and calls the underlying driver QueryContext method if it is supported.
func (d *EntMysqlDriver) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	drv, ok := d.Driver.(interface {
		QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	start := time.Now()
	rows, err := drv.QueryContext(ctx, query, args...)
	duration := time.Since(start)
	d.startTrace(ctx, query, duration, args)
	d.log(ctx, fmt.Sprintf("driver.QueryContext: query=%v args=%v time=%v", query, args, duration))
	return rows, err
}

func (d *EntMysqlDriver) startTrace(ctx context.Context,
	query string, duration time.Duration, args ...any) {
	if d.isTrace {
		span := trace.SpanFromContext(ctx)
		if span.IsRecording() {
			attrs := []attribute.KeyValue{
				mysqlQueryTextKey.String(query),
				mysqlArgsTextKey.String(fmt.Sprint(args)),
				mysqlDurationKey.String(fmt.Sprintf("%v", duration)),
			}
			span.AddEvent(mysqlEventKey, trace.WithAttributes(attrs...))
		}
	}
}

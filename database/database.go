package database

import (
	"database/sql"
	"fmt"
	t "time"

	_ "github.com/lib/pq"
	"k8s.io/klog"
)

type Database struct {
	db             *sql.DB
	preparedInsert *sql.Stmt
	preparedDelete *sql.Stmt
}

// NewDatabase will crate a new database connection
func NewDatabase(dbname, user, host, password, sslmode, schema string, port int, timescale bool) (Database, error) {
	conStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	d, err := sql.Open("postgres", conStr)
	if err != nil {
		return Database{}, err
	}

	data := Database{
		db: d,
	}

	if err = data.crateTable(schema, timescale); err != nil {
		return Database{}, err
	}

	in := fmt.Sprintf("INSERT INTO %s.devices (time, value, device, namespace, sensor, active) VALUES ($1, $2, $3, $4, $5, $6)", schema)
	prepInsert, err := d.Prepare(in)
	if err != nil {
		errorString := fmt.Sprintf("could not prepare insert stm %v; err: %v\n", in, err)
		if er := d.Close(); er != nil {
			return Database{}, fmt.Errorf("%s \n\t and could not close databse successfully; err: %v\n", errorString, er)
		}
		return Database{}, fmt.Errorf("%s\n", errorString)
	}
	prepDelete, err := d.Prepare(fmt.Sprintf("UPDATE %s.devices SET active = 'false' WHERE namespace = $1 AND device = $2", schema))
	if err != nil {
		errorString := fmt.Sprintf("could not prepare delete; err: %v\n", err)
		if er := d.Close(); er != nil {
			return Database{}, fmt.Errorf("%s \n\t and could not close databse successfully; err: %v\n", errorString, er)
		}
		return Database{}, fmt.Errorf("%s\n", errorString)
	}

	data = Database{
		db:             d,
		preparedInsert: prepInsert,
		preparedDelete: prepDelete,
	}

	return data, nil
}

// Close will close the database connection
func (d Database) Close() error {
	return d.db.Close()
}

// Insert will insert data into a database table
func (d Database) Insert(time t.Time, value, device, namespace, sensor string, active bool) error {
	_, err := d.preparedInsert.Exec(time, value, device, namespace, sensor, active)
	return err
}

// Delete will delete a database table
func (d Database) Delete(device, namespace string) error {
	_, err := d.preparedDelete.Exec(namespace, device)
	return err
}

func (d Database) crateTable(schema string, timescale bool) error {
	var stm string
	if schema == "" {
		schema = "public"
	}
	stm = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.devices (time TIMESTAMP, value TEXT, device TEXT, namespace TEXT, sensor TEXT, active BOOL)", schema)
	if _, err := d.db.Exec(stm); err != nil {
		klog.Errorf("could not crate table %v\n", stm)
		return err
	}
	if timescale {
		tableExists, err := d.db.Query(fmt.Sprintf("SET search_path TO %s; SELECT exists ( SELECT * FROM timescaledb_information.hypertable WHERE table_name = 'devices')", schema))
		if err != nil {
			return err
		}

		var exists bool
		for tableExists.Next() {
			if err = tableExists.Scan(&exists); err != nil {
				_ = tableExists.Close()
				return err
			}
		}

		if err = tableExists.Close(); err != nil {
			return err
		}

		if !exists {
			if _, err := d.db.Exec(fmt.Sprintf("SET search_path TO %s; SELECT * FROM CREATE_HYPERTABLE('devices', 'time', number_partitions => 4)", schema)); err != nil {
				return err
			}
			if _, err := d.db.Exec(fmt.Sprintf("SET search_path TO %s; SELECT * FROM add_dimension('devices', 'namespace', number_partitions => 4)", schema)); err != nil {
				return err
			}
			if _, err := d.db.Exec(fmt.Sprintf("SET search_path TO %s; SELECT * FROM add_dimension('devices', 'device', number_partitions => 4)", schema)); err != nil {
				return err
			}
			if _, err := d.db.Exec(fmt.Sprintf("SET search_path TO %s; SELECT * FROM add_dimension('devices', 'sensor', number_partitions => 4)", schema)); err != nil {
				return err
			}
		}
	}
	return nil
}

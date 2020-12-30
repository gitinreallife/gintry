package models

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
)

// "github.com/satori/go.uuid"

type LASTUser struct {
	gorm.Model
	nomor         string
	nama          string
	username      string
	password      string
	fidrole       string
	kodeCabang    string
	limit         float32
	fidJhooff     string
	status        string
	sbdescUker    string
	unitKerjaUker string
	limitKupedes  float64
	limitKur      float64
	limitBriguna  float64
	limitRitel    float64
	limitPangan   float64
	limitMenengah float64
	limitCashcoll float64
}

func ReadLASTUsers(where string) (int, error) {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}
	tsql := fmt.Sprintf("Select top 1 UID, NO_PEGAWAI, NAMA_PEGAWAI, USER_NAME from LAS_T_USER where 1=1")
	if where != "" {
		tsql += fmt.Sprintf(" and %s;", where)
	}
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	var count int
	var noPegawai, namaPegawai, userName string
	var uid int
	for rows.Next() {
		err := rows.Scan(&uid, &noPegawai, &namaPegawai, &userName)
		if err != nil {
			return -1, err
		}
		fmt.Printf("UID: %d, NO_PEGAWAI: %s, NAMA_PEGAWAI: %s, USER_NAME: %s\n",
			uid, noPegawai, namaPegawai, userName)
		count++
	}
	return count, nil
}

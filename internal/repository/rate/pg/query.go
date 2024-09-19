package pg

const sqlRAInsertQuery string = `
     insert into rate_agroups values ($1, $2, $3, $4, $5, $6, $7)`

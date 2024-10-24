package oracle

const sqlRACountQuery string = `
     select count(*)
       from rms_rate_agroup_codes ac
       join rms_rate_agroups       a on a.rmsg_id   = ac.rmsg_id
       join rms_rate_bgroups       b on b.dest_name = a.dest_name
      where ac.dend > sysdate
        and ac.direction = %s`

const sqlRAQuery string = `
     select /*+ index(ac PK_RMSRAGC) */ b.rmsg_id as b_rmsg_id
            ,ac.rmsg_id as a_rmsg_id
            ,ac.gwgr_id
            ,ac.direction
            ,ac.dial_code
            ,ac.dbegin
            ,ac.dend
        from rms_rate_agroup_codes ac
        join rms_rate_agroups       a on a.rmsg_id   = ac.rmsg_id
        join rms_rate_bgroups       b on b.dest_name = a.dest_name
       where ac.dend > sysdate
         and ac.direction = %s`

const sqlRBCountQuery string = `
	 select count(*)
       from rms_rate_bgroup_codes bc
      where bc.dend > sysdate
        and bc.direction = %s`

const sqlRBQuery string = `
	select /*+ index(bc PK_RMSRBGC) */ bc.rmsg_id
          ,bc.gwgr_id
          ,bc.direction
          ,bc.dial_code
          ,bc.dbegin
          ,bc.dend
      from rms_rate_bgroup_codes bc
     where bc.dend > sysdate
       and bc.direction = %s`

const sqlRTSCountQuery string = `
    select count(*)
	  from rms_rates r
	 where r.dend > sysdate`

const sqlRTSQuery string = `
   select r.gwgr_id
         ,r.direction
         ,r.a_rmsg_id
         ,r.b_rmsg_id
         ,r.rmsr_id
         ,r.rmsv_id
         ,r.dbegin
         ,r.dend
     from rms_rates r
    where r.dend > sysdate`

const sqlRVCountQuery string = `
     select count(*)
       from rms_rate_values v`

const sqlRVQuery string = `
     select v.rmsv_id
           ,v.currency_id
           ,nvl(v.price1, 0) as price
       from rms_rate_values v`

const sqlCURRTSCountQuery string = `
     select count(*)
       from currency_rates r
      where r.dend > sysdate`

const sqlCURRTSQuery string = `
     select r.currency_id
           ,r.rate
           ,r.dbegin
           ,r.dend
       from currency_rates r
      where r.dend > sysdate`

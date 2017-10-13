select current_database();
select version();
select current_time;
select inet_server_addr();
select inet_server_port();
select pg_postmaster_start_time();
select date_trunc(‘minute’, current_timestamp – pg_postmaster_start_time()) as “postgresql uptime”;
select pg_database_size(‘uptimemadeeasy’);
select pg_total_relation_size(‘awards’);
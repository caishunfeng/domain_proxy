内网域名代理服务

添加域名代理：
curl --location --request GET 'http://proxy.local/proxy/add?domain=csf.proxy.local&port=8081&expire=24'
domain: 想代理的内网域名,格式必须是*.xxx,如*.proxy.local
port: 想代理的本地端口
expire: 有效时间（小时），默认为12小时

nginx配置：
    server {
        listen   8080;
        server_name   proxy.local; // your host name
        access_log logs/proxy_$year$month$day.log main;

        location / {
          proxy_set_header Host $host;
          proxy_set_header X-Real_IP $remote_addr;
          proxy_set_header X-Real-Port $remote_port;
          proxy_set_header X-Timestamp $msec;
          proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
          proxy_pass http://intranet_ip:8082/;
        }
    }

    server {
        listen   8080;
        server_name   *.proxy.local; // your host name
        access_log logs/proxy_$year$month$day.log main;

        location / {
          proxy_set_header Host $host;
          proxy_set_header X-Real_IP $remote_addr;
          proxy_set_header X-Real-Port $remote_port;
          proxy_set_header X-Timestamp $msec;
          proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
          proxy_pass http://intranet_ip:8082/;
        }
    }
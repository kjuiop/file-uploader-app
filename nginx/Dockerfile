FROM nginx:1.27.0

WORKDIR /etc/nginx

COPY nginx.conf /etc/nginx/nginx.conf
COPY conf.d /etc/nginx/conf.d

# 로그 디렉터리 생성 및 권한 설정
RUN mkdir -p /var/log/nginx \
    && touch /var/log/nginx/kus.access.log \
    && touch /var/log/nginx/kus.error.log \
    && chmod 644 /var/log/nginx/*.log \
    && mkdir -p /etc/nginx/ssl \
    && chmod 644 /etc/nginx/ssl \
    && mkdir -p /run/nginx \
    && chmod 755 /run/nginx

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]

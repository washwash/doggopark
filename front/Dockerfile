FROM nginx:latest as build

COPY app /opt/app
COPY resources/nginx.conf /etc/nginx/conf.d/
RUN service nginx restart

EXPOSE 80

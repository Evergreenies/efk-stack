FROM fluent/fluentd:v1.16-debian-1
USER root
RUN gem install fluent-plugin-elasticsearch
RUN gem install fluent-plugin-concat 

USER fluent 
COPY ./conf/fluent.conf /fluentd/etc/fluent.conf

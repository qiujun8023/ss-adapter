FROM alpine

ENV APP_ROOT /app
ENV NODE_ENV production

WORKDIR ${APP_ROOT}

RUN echo -e "https://mirrors.ustc.edu.cn/alpine/latest-stable/main\nhttps://mirrors.ustc.edu.cn/alpine/latest-stable/community" > /etc/apk/repositories && \
    apk update && apk --no-cache add jq

COPY ss-adapter entrypoint.sh ${APP_ROOT}/

CMD ["sh", "entrypoint.sh"]

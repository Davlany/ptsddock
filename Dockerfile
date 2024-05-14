FROM alpinelinux/golang

WORKDIR /dock

COPY . .

CMD ["go", "test", "-v"]
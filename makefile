OUT_PATH = "out/whow"

build:
	go build -o ${OUT_PATH} cmd/whow

run: build
	${OUT_PATH}

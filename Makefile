doc:
	swag init -g ./cmd/main.go -o docs && mv docs/swagger.yaml cmd/swagger.yml

clean:
	rm -rf docs

all:
	make doc && make clean

.PHONY: doc clean all

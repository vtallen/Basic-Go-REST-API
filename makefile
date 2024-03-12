TARGET := restapi
BUILDDIR := bin

all: swaggervalid swaggerdoc 
	mkdir -p $(BUILDDIR)
	go build -o $(BUILDDIR)/$(TARGET) internal/*

swaggergen:
	go generate github.com/vtallen/Basic-Go-REST-API/internal github.com/vtallen/Basic-Go-REST-API/pkg/swagger

swaggervalid:
	swagger validate pkg/swagger/swagger.yml

swaggerdoc:
	sudo docker run -i yousan/swagger-yaml-to-html < pkg/swagger/swagger.yml > doc/index.html

clean:
	rm -r $(BUILDDIR)
	rm -r pkg/swagger/server

run:
	./$(BUILDDIR)/$(TARGET)

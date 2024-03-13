TARGET := restapi
BUILDDIR := bin

all: swaggervalid swaggerdoc swaggergen build
	
build:
	mkdir -p $(BUILDDIR)
	go build -o $(BUILDDIR)/$(TARGET) internal/*

swaggergen:
	@echo "Generating swagger server"
	@echo "========================="
	@echo ""
	go generate github.com/vtallen/Basic-Go-REST-API/internal github.com/vtallen/Basic-Go-REST-API/pkg/swagger
	@echo ""

swaggervalid:
	@echo "Validating swagger.yml"
	@echo "======================="
	@echo ""
	swagger validate pkg/swagger/swagger.yml
	@echo ""

swaggerdoc:
	@echo "Generating swagger documentation"
	@echo "================================"
	@echo ""
	sudo docker run -i yousan/swagger-yaml-to-html < pkg/swagger/swagger.yml > doc/index.html
	@echo ""

clean:
	rm -r -f $(BUILDDIR)
	rm -r -f pkg/swagger/server

run:
	./$(BUILDDIR)/$(TARGET)

.PHONY : clean

golab:
	@echo "start of golab"
	go build -o golab main.go

clean :
	@echo "start of clean"
	-rm golab main
##
run:
	./golab

echo:
	@ls ~/sql
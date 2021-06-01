.PHONY : all

README.md: add.md
	@echo "add发生了变化"
add.md:add2.md
	@echo "add2发生了变化"

all: part1 part2 part3

part1:
	@echo "i am part 1\n"

part2:
	@echo "i am part 2\n"

part3:
	@echo "i am part 3\n"


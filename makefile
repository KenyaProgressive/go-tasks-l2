task1:
	go run 1/main.go

task2:
	go run 2/main.go

task3:
	go run 3/main.go

task4:
	go run 4/main.go

task5:
	go run 5/main.go

task6:
	go run 6/main.go

task7:
	go run 7/main.go

task8:
	cd task8-ntp && go run . && cd ..

task8-pretty-check:
	golint task8-ntp/
	cd task8-ntp && go vet task8-ntp/ && cd ..
	goimports -d task8-ntp/
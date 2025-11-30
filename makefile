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

task9:
	go run 9/main.go

task9-pretty-check:
	golint 9/
	cd 9 && go vet 9/ && cd ..
	goimports -d 9/
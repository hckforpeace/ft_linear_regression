
ALL: predict precision train 

predict:
	go build -o predict  ./cmd/predict/main.go 

precision:
	go build -o precision  ./cmd/precision/main.go 

train:
	go build -o train  ./cmd/train/main.go 

clean:
	rm precision train predict

re: clean ALL

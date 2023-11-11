brew install protoc-gen-go /
brew install protobuf 

protoc --go_out=pkg/api --go_opt=paths=source_relative \
--go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative \
proto/adder.proto

https://www.youtube.com/watch?v=z-mHhobE0Pw&t=868s

protoc - компилятор пртобуфа
-I=pacage - что у нас является корнем, где искать файлы
--go_out=pkg/api - какой файл собрем и куда ложить сгенерированные файлы

buf.yaml - конфиги для линтера, конфиг для ломаюших изменений для протобуффа - единая система для конфигурации протобуфа
buf.gen.yaml - указываем плагины куда сгенерировать, опции которые нам нужны

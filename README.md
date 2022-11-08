# test-rover

#### Para executar o programa é necessário ter o ambiente Docker instalado.

### Execute o programa com o arquivo data/input.txt + testes:
```
make
```

### Também é possível executar diretamente na máquina:
```
go run cmd/main.go data/input.txt
```

### Você pode verificar se o test-rover funciona apenas executando os testes:
```
go test ./...
```

### Posição e direção
Para direcionar o rover para esquerda (-1), direita (+1) e mover (0), usamos uma operação de soma. Considerando que norte, leste, sul, oeste são representados respectivamente por 0, 1, 2, 3. Assuminos também que -1 equivale a oeste e 4 a norte.

### Mover
Para mover o rover na grande, consideramos o ponto X como sendo oeste (-1) e leste (+1). E o ponto Y sendo norte (+1) e sul (-1).
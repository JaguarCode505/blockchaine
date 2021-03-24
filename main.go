package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC to Daria")
	bc.AddBlock("Send 2 more BTC to Valihan")
	for _, block :=range bc.blocks{
		prevHash := block.prevBlockHash
		fmt.Println(prevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}

type Block struct { //создаем структуру каждого из блоков которые и представляют из себя blockchain
	Timestamp int64 //Временная метка когда да
	Data []byte //Биткоины
	Hash []byte // идентификатор SHA256, идентифицирующий текущую запись
	prevBlockHash []byte //идентификатор SHA256, идентифицирующий прошлую запись в цепочке
}
/*
	Мы используем хешировние для сохранение целостности цепочки.
	Сохраняя предыдущие хэши, как мы делаем на диаграмме выше, мы можем гарантировать, что блоки в blockchain находятся в правильном порядке.
	Если злоумышленник захочет присоединитьсяФ и манипулировать данными (например, изменить количество денег в своем кошелке,
	хэши начнут изменяться и все будут знать, что цепочка "сломана" и все будут знать, что доверять этой цепочки нельзя.
*/
func (b*Block) setHash(){ ////Создаем функцию которая соединяет все данные с блока(для упрощения хеширования) и создает Hash функцию SHA256
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.prevBlockHash,b.Data,timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block{ //создаем конструктор для блока
	block:= &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.setHash()
	return block
}

func NewGenesisBlock() *Block { //функция создающая genesis блок
	return NewBlock("Genesis Block", []byte{})
}

type Blockchain struct { //блокчейн
	blocks []*Block
}


func (bc *Blockchain) AddBlock(data string) { //Возможность добавлять блоки в блокчейн
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
/*
	(Костыль?)Чтобы добавить первый блок,нужен уже существующий, но  блокчейн пуст в самом начале
	Поэтому в любом блокчейне должен быть как минимум один блок, который называют genesis блоком.
*/

func NewBlockchain() *Blockchain { //функция создающая блокчейн с genesis блоком
	return &Blockchain{[] *Block{NewGenesisBlock()}}
}
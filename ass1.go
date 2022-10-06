package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type Block struct {
	blockno int
	trans   string
	nonce   int
	phash   string
	hash    string
}

func NewBlock(num int, tran string, nonc int, ph string, ch string) *Block {
	s := new(Block)
	s.blockno = num
	s.trans = tran
	s.nonce = nonc
	s.phash = ph
	s.hash = ch
	return s
}

type BlockChain struct {
	list []*Block
}

func (ls *BlockChain) CreateBlock(tran string, nonc int) *Block {
	//Retriving old hash
	length := len(ls.list)
	ph := ""
	if length != 0 {
		ph = ls.list[length-1].hash
	}

	//Calculating current hash
	nonce2string := strconv.Itoa(nonc)
	num2string := strconv.Itoa(length)
	data := num2string + tran + nonce2string + ph
	var c_hash string = CalculateHash(data)

	//Creating New Block
	st := NewBlock(length, tran, nonc, ph, c_hash)
	ls.list = append(ls.list, st)
	return st
}

func (ls *BlockChain) ListBlocks() {
	length := len(ls.list)
	i := 0
	fmt.Println("\n**********************************")
	for i < length {
		fmt.Println("\n\n\tBlock Number = ", ls.list[i].blockno)
		fmt.Println("\tTransection  = ", ls.list[i].trans)
		fmt.Println("\tNonce        = ", ls.list[i].nonce)
		fmt.Println("\tPrev Hash    = ", ls.list[i].phash)
		fmt.Println("\tHash         = ", ls.list[i].hash)
		i = i + 1
	}
	fmt.Println("\n==================================================================================================")

}

func CalculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func (ls *BlockChain) ChangeBlock(num int, tran string) {
	ls.list[num].trans = tran
}

func (ls *BlockChain) VerifyChain() {
	length := len(ls.list)
	i := 0
	var check bool = true
	index := 0
	for i < length-1 {

		//Current Block Data and Nxt blocks Prev Hash
		C_nonce2string := strconv.Itoa(ls.list[i].nonce)
		C_num2string := strconv.Itoa(ls.list[i].blockno)
		C_data := C_num2string + ls.list[i].trans + C_nonce2string + ls.list[i].phash
		var C_h string = CalculateHash(C_data)

		var Nxt_h string = ls.list[i+1].phash

		fmt.Println("\n\n\t**VALIDITY TEST ", i, "*")
		fmt.Println("\n\t Nxt_Block_Prev_Hash  = ", Nxt_h)
		fmt.Println("\n\t Curr_Block_Hash      = ", C_h)

		if Nxt_h != C_h {
			check = false
			index = i
			break
		}
		i++
	}

	if i == length-1 && check == true && index != i {
		//Current Block Data and Hash
		C_nonce2string := strconv.Itoa(ls.list[i].nonce)
		C_num2string := strconv.Itoa(ls.list[i].blockno)
		C_data := C_num2string + ls.list[i].trans + C_nonce2string + ls.list[i].phash
		var C_h string = CalculateHash(C_data)
		var SC_h string = ls.list[i].hash
		fmt.Println("\n\n\t**VALIDITY TEST ", i, "*")
		fmt.Println("\n\t Stored_Curr_Hash = ", SC_h)
		fmt.Println("\n\t Curr_Hash        = ", C_h)
		if SC_h != C_h {
			check = false
			index = i
		}
	}

	fmt.Println("\n\n\t**VALIDITY TEST*")
	if check {
		fmt.Println("\n\tChain Is Valid!")
	} else {
		fmt.Println("\n\tChain Is Not Valid!\n\tError in Block Number ", i)
	}

}

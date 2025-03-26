package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

// Contador de IPs em uso e Mutex para sincronização
var (
	inUseCount int
	mutex      sync.Mutex
)

func ping(ip string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Contador de IPs em uso e Mutex para sincronização
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("ping", "-n", "1", "-w", "100", ip) // -n para Windows
	} else {
		cmd = exec.Command("ping", "-c", "1", "-W", "1", ip) // -c para Linux
	}

	// Executa o ping e verifica o resultado
	if err := cmd.Run(); err != nil { //Imprime e conta se IP estiver disponiveis
		mutex.Lock()
		inUseCount++
		mutex.Unlock()
		fmt.Printf("%s está disponível\n", ip)
	}
}

/*
	} else {
		fmt.Printf("%s está em uso\n", ip)
	}
*/

func main() {
	baseIP := "172.19.1"
	var wg sync.WaitGroup

	for i := 1; i <= 254; i++ {
		ip := fmt.Sprintf("%s.%d", baseIP, i)
		wg.Add(1)
		go ping(ip, &wg)
		time.Sleep(10 * time.Millisecond) // Delay
	}

	wg.Wait()
	fmt.Println("Verificação concluída.\n Numero total de IP's disponíveis:", inUseCount)
	//	fmt.Println("Verificação concluída.\n Numero total de IP's em uso:", inUseCount)
}

package main

import "fmt"

func main() {
	// definimos y almacenamos valores
	var CloudProviders = [4]string{"Google Cloud", "Azure", "AWS", "Huawei(sketchy)"}

	for i := 0; i < len(CloudProviders); i++ {
		option := CloudProviders[i]
		fmt.Println(i, option)
	}

	// Ahora con un for loop more "cheto"
	RealCloudProviders := [...]string{"GC", "AWS", "Azure"}

	fmt.Println("\nTha Real Ones")
	for index, option := range RealCloudProviders {
		fmt.Println(index, option)
	}

}

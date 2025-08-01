package main

import "fmt"

func main() {
	fmt.Println("maps in golang")

	// 1️⃣ Map banana: string keys and string values
	languages := make(map[string]string)

	// 2️⃣ Values add kar rahe hain map me
	languages["r"] = "ruby"
	languages["c"] = "cpp"

	// 3️⃣ Pura map print
	fmt.Println("list all", languages)

	// 4️⃣ Specific key ka value print
	fmt.Println("r is", languages["r"])

	// 5️⃣ Key "r" ko hata do
	delete(languages, "r")

	// 6️⃣ Dobara map print (ab "r" nahi hoga)
	fmt.Println("list all", languages)

	// 7️⃣ Map ke upar loop — key aur value dono
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

	// 8️⃣ Agar sirf value chahiye ho to:
	for _, value := range languages {
		fmt.Printf("Only value: %v\n", value)
	}
}

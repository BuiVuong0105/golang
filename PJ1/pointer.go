// package main

// import "fmt"

// type Player struct {
// 	health int
// }

// // định nghĩa hàm takeDamageFromExplosion trong struct player, sử dụng con trỏ Player chứ ko sử dụng đối tượng
// // Hàm này nhận player là 1 tham số ngầm định
// func (player *Player) takeDamageFromExplosion(damage int) {
// 	fmt.Println("Player is taking damage from explosion")
// 	player.health -= damage
// }

// // takeDamageFromExplosion nhận param là con trỏ Player sẽ tương tự với hàm ở trên.
// func takeDamageFromExplosion(player *Player, damage int) {
// 	fmt.Println("Player is taking damage from explosion")
// 	player.health -= damage
// }

// func main() {

// 	player := &Player{
// 		health: 100,
// 	} // 8 byte long integer (pointer)
// 	fmt.Printf("Before explosion: %+v\n", player)
// 	// player = nil // somewhere in the game you deleted player
// 	player.takeDamageFromExplosion(10)
// 	takeDamageFromExplosion(player, 10)
// 	fmt.Printf("After explosion: %+v\n", player)
// }

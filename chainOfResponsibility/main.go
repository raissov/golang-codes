package main

func main()  {
	gameWindow:=&GamePage{}

	heroWindow:=&HeroPage{}
	heroWindow.setNext(gameWindow)

	mainWindow:=&mainPage{}
	mainWindow.setNext(heroWindow)

	zadrot:=&player{nickName:"Beka"}
	mainWindow.execute(zadrot)
}

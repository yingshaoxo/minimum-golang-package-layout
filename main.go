package main

import (
	"github.com/yingshaoxo/my_package/another_sub_module"
	final_sub_module "github.com/yingshaoxo/my_package/final_module"
	sub_module "github.com/yingshaoxo/my_package/sub_module_folder"
)

func main() {
	sub_module.Print_something()
	another_sub_module.Print_something_else()
	final_sub_module.Print_the_final_words()
}

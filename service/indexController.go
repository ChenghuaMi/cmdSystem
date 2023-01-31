/**
 * @author mch
 */

package service

import "fmt"

type  IndexController struct {

}
func(c *IndexController) Welcome() {
	fmt.Println("welcome to....")
	view = "login_view"

}
func(c *IndexController) Index() {
	fmt.Println("come to index")
	view = "index_view"

}

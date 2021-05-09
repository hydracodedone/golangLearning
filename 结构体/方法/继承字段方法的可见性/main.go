package main

import somepkg "demo/somePkg"

func main() {
	var public somepkg.Public
	public.PrivateHasPublicMethod() //可以访问,匿名字段,可以直接调用该方法,但是不能通过public.private调用,因为private字段小写,其他包不可见
	// public.publicHasPrivateMethod() //不可访问
}

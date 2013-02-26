class.page
==========

a go class for page 

找半天分页类，没有？自己写一个吧

#用法：


` //获取分页变量
page := util.NewPageMaker(num, p, "list")
//获取用户信息
userList := user.list(page)

func (u *User) list(page *util.PageMaker) []UserData {
    //fmt.Println("user.go")
    //判断是否重复
    var l []UserData
    u.c.Find(nil).Sort("-_id").Skip(page.Begin).Limit(page.PerPage).All(&l)
   
    return l
   
}`


#显示：
第一页 << 11 12 13 14 15 >> 最后页

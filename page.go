// 获取分页
// 找半天分页类，没有？自己写一个吧
// by indraw 2013.02.01
package util

import (
  "math"
    "fmt"
	"html/template"
)
//定义用户信息
type PageMaker struct {
    Page     int       //当前页
    PageMax  int       //最大分页
    PagePre  int       //上一页
    PageNext int       //下一页
    Count    int       //总数
    PerDiv   int       //每段数量
    PerPage  int       //每页数量
    
    Url      string    //当前链接
    List     template.HTML    //页码字符串
    Begin    int       //开始数量
    Top      int       //每页最大ID
    
    
}

/**
 * 生成分页
 *
 * author: indraw
 * date: 2013-02-01
 */
func  NewPageMaker( count int,  p int , url string,) *PageMaker {
    //初始化
    page :=  &PageMaker{
                Page:1,
                PageMax:1,
                PagePre:1,
                PageNext:1,
                Count:0,
                PerDiv:10,
                PerPage:5,
                Url:url,
                List:"",
                Begin:0,
                Top:0,
            }
    var html string
    page.Page = p
    if page.Page < 1 {
        page.Page = 1
    }
    //计算分页数据
    page.Count = count
    page.Url = url

    page.PageMax = int(math.Ceil( float64(page.Count) / float64(page.PerPage) ) )
    page.Begin = page.PerPage * (page.Page -1 )
    
    //获取第一页
    if page.Page > page.PerPage{
        if page.Page == 1{
            html = ""
        }else{
            html = "<a href='" + page.Url + "' title='首页'>1</a>\n"
            
        }
    }
    //循环获取分页信息
    var f int = 1
    var l int = page.PerPage
    for f = 1; f <= page.PageMax; f++ {
        if page.Page >= f && page.Page <= l {
            
            //计算前页
            page.PagePre = f - 1
            if page.PagePre > 0 {
                html = fmt.Sprintf("%s<a href=\"%s?p=%d\" title=\"前 %d 页\">Prev</a>-\n", html, page.Url, page.PagePre, page.PerDiv)
            }else{
                html = html + "\n"
            }
            
            //生成分页列表
            if l <= page.PageMax{
                for j := f; j <= l; j++ {
                    if j == page.Page {
                        html = fmt.Sprintf("%s<a href=\"%s?p=%d\"  class=\"on\" title=\"当前页\">%d</a>\n", html, page.Url, page.Page, j)
                    }else{
                        if j == 1{
                            html = fmt.Sprintf("%s<a href=\"%s\" title=\"第1页\">1</a>\n", html, page.Url)
                        }else{
                            html = fmt.Sprintf("%s<a href=\"%s?p=%d\" title=\"第%d页\">%d</a>\n", html, page.Url, j, j, j)
                        }
                    }
                }
                
            }else{
                for j := f; j <= page.PageMax; j++ {
                    if j == page.Page{
                        html = fmt.Sprintf("%s<a href=\"%s?p=%d\"  class=\"on\" title=\"当前页\">%d</a>\n", html, page.Url, page.Page, j)
                    }else{
                        if j == 1{
                            html = fmt.Sprintf("%s<a href=\"%s\" title=\"第1页\">1</a>\n", html, page.Url)
                        }else{
                            html = fmt.Sprintf("%s<a href=\"%s?p=%d\" title=\"第%d页\">%d</a>\n", html, page.Url, j, j, j)
                        }
                    }
                }
                
            }
            
            //计算后页

            page.PageNext = l + 1
            if page.PageNext < page.PageMax{
                html = fmt.Sprintf("%s-<a href=\"?p=%d\" title=\"后 %d 页\">Next</a>\n", html, page.PageNext, page.PerDiv)

            }else{
                html = html + "\n"
            }
        }
        
        f = f + page.PerPage
        l = l + page.PerPage
    }
    
    //
    if page.PageNext <= page.PageMax{
        html =  fmt.Sprintf("%s<a href='?p=%d' title='尾页'>%d</a>\n", html, page.Page, page.PageMax)
    }
    
    page.List = template.HTML(html)
    
    page.Top = page.Count - page.PerPage*(page.Page - 1)
    
    return page
}


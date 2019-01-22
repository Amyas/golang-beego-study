{{define "nav"}}
<div class="navbar navbar-default navbar-fixed-top">
  <div class="container">
    
    <a class="navbar-brand" href="/">我的博客</a>
    
    <ul class="nav navbar-nav">
      <li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
      <li {{if .IsCategory}}class="active"{{end}}><a href="/category">分类</a></li>
      <li {{if .IsTopic}}class="active"{{end}}><a href="/topic">文章</a></li>
    </ul>

    <div class="pull-right">
      <ul class="nav navbar-nav">
        <li>
          {{if .IsLogin}}
            <a href="/login?exit=true">退出</a>
          {{else}}
            <a href="/login">管理员登录</a>
          {{end}}
        </li>
      </ul>
    </div>
    
  </div>
</div>
{{end}}
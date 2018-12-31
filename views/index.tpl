<<template "header.html" .>>
<<template "navbar.html" .>>

<section role="main" class="main index">
<ul class="posts">
  << range $v := .posts>>
  <li>
      <article class="post">
        <header>
          <h2 class="page-header"><a data-pjax="" href="/post/<< $v.File >>"><< $v.Title >></a></h2>
        </header>
        <section class="post-content">
          << $v.Date >>
          <br />
          << str2html $v.Summary >>
        </section>
    </article>
  </li>
  << end >>
</ul>

<<template "footer.html" .>>

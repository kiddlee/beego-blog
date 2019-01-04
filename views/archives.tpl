<<template "header.html" .>>
<<template "navbar.html" .>>
<section role="main" class="main page">
	<article class="post">
  		<header>
    		<h2 class="page-header"><a data-pjax="" href="/archives">Archives</a></h2>
  		</header>
		<section class="post-content">
            <ul>
                << range $v := .posts >>
                <li><a href="/post/<< $v.File >>"><< $v.Title >></a></li>
                << end >>
            </ul>
        </section>
    </article>
</section>
<<template "footer.html" .>>

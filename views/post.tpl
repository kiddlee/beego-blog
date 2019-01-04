<<template "header.html" .>>
<<template "navbar.html" .>>
<section role="main" class="main page">
	<article class="post">
        <header>
	    	<h2 class="page-header"><a data-pjax="" href="/post/<< .post.File >>"> << .post.Title >></a></h2>
	          	<div class="meta">
	                <ul class="tags">
                        << range $v := .post.Tags >>
	                    <li><a class="tag" href="/tag/<< $v >>"><< $v >></a></li>
                        << end >>
	                </ul>
	            </div>
	  	</header>
		<section class="post-content">
          	<< str2html .post.Body >>
      	</section>
        <footer>
	      	<section class="meta-date">
	        	posted on
	        	<time datetime="<< .post.Date >>" pubdate="">
	          		<span><< .post.Date >></span>
	        	</time>
	      	</section>
	    </footer>
    </article>
</section>
<<template "footer.html" .>>

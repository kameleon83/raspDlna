<h1 class="text-center">{{ .title }}</h1>
<br/>

<div class="container">
    <form class="col-md-6 col-md-offset-3" action="/login" method="post">
        <div class="form-group">
            <input type="text" class="form-control input-lg username" placeholder="Name" name="username">
        </div>
        <div class="form-group">
            <input type="password" class="form-control input-lg password" placeholder="Password" name="password">
        </div>
        <div class="form-group">
            <button class="btn btn-primary btn-lg btn-block">Se connecter</button>
        </div>
    </form>
    <h4 class="col-md-6 col-md-offset-3 text-center">Si cela est la première fois que tu es sur ce site il faut en configurer l'application.
        <br/><br/>
        <a href="/register">Cliques ici pour configurer l'application.</a>
    </h4>
</div>

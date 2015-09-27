<h1 class="text-center">{{ .title }}</h1>

<div class="container">
    <form class="col-md-6 col-md-offset-3" action="/register" method="post">
        <div class="form-group">
            <input type="name" class="form-control input-lg" placeholder="Name" name="name">
        </div>
        <div class="form-group">
            <input type="password" class="form-control input-lg" placeholder="Password" name="password">
        </div>
        <div class="form-group">
            <input type="password" class="form-control input-lg" placeholder="Confirm Password" name="confirmPassword">
        </div>
        <div class="form-group">
            <input type="text" class="form-control input-lg" placeholder="Path Folders" name="root" value={{ .root }}>
        </div>
        <div class="form-group">
            <button class="btn btn-primary btn-lg btn-block">Enregistrer la configuration</button>
        </div>
    </form>
</div>

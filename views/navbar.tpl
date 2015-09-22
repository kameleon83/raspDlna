<nav class="navbar navbar-default" role="navigation">
    <div class="container-fluid">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#navbar">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#"></a>
        </div>

        <!-- Collect the nav links, forms, and other content for toggling -->
        <div class="collapse navbar-collapse" id="navbar">

            <div class="navbar-form navbar-left">
                <button type="button" class="btn btn-danger" data-toggle="modal" data-target="#mkdir">
                    Créer un nouveau dossier
                </button>
            </div>
            <div class="navbar-form navbar-left">
                <button type="button" class="btn btn-danger" data-toggle="modal" data-target="#chown">
                    Changer les droits
                </button>
            </div>
            <form action="/cmdperso{{ .href }}" method="post">
            <div class="navbar-form navbar-right">
                <input type="input" class="form-control" placeholder="Cmd Perso : " size="50" name="cmdperso" value="" />
                <button type="submit" class="btn btn-primary">
                    Valider
                </button>
            </div>
        </form>
        </div><!-- /.navbar-collapse -->
    </div><!-- /.container-fluid -->
</nav>



<div class="modal fade" id="mkdir" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
    <div class="modal-dialog" role="document">
        <form method="POST" class="form-horizontal col-sm-12" action="/mkdir{{ .href }}">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title" id="myModalLabel">Créer un dossier ici</h4>
                </div>
                <div class="modal-body">
                        <fieldset>
                            <div class="col-sm-10 form-group">
                                <input type="text" class="form-control" id="mkdir"  placeholder="Créer un nouveau dossier" name="mkdir" value="" />
                            </div>
                        </fieldset>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                    <button type="submit" class="btn btn-primary">
                        Enregistré la modification !
                    </button>
                </div>
            </div>
        </form>
    </div>
</div>


<div class="modal fade" id="chown" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
    <div class="modal-dialog" role="document">
        <form method="POST" class="form-horizontal" action="/chown">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title" id="myModalLabel">Changer utilisateur sur tous les fichiers et dossiers</h4>
                </div>
                <div class="modal-body">

                    <fieldset>
                        <div class="col-md-10 form-group">
                            <label for="listUid">Choisir Utilisateur</label>
                            <select class="form-control" name="user" id="listUid">
                                {{ range $k, $v := .user }}
                                <option value="{{ $v }}">{{ $k }}</option>
                                {{ end }}
                            </select>
                        </div>
                        <div class="col-md-10 form-group">
                            <label for="listGid">Choisir Groupe</label>
                            <select class="form-control" name="group" id="listGid">
                                {{ range $k, $v := .group }}
                                <option value="{{ $v }}">{{ $k }}</option>
                                {{ end }}
                            </select>
                        </div>
                        <br>
                        <br>
                        <br>
                        <div class="col-md-10 form-group">
                            <label for="pass">Rentrer le mot de passe Sudo</label>
                            <input type="password" class="form-control" id="pass"  placeholder="Sudo Password" name="pass" value="" />
                        </div>
                    </fieldset>
                </div>
                <div class="modal-footer">

                    <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                    <button type="submit" class="btn btn-primary">
                        Enregistré la modification !
                    </button>
                </div>
            </div>
        </form>
    </div>
</div>

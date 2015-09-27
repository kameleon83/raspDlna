<nav class="navbar navbar-fixed-top navbar-default" role="navigation">
    <div class="container-fluid">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#navbar">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>

        </div>

        <!-- Collect the nav links, forms, and other content for toggling -->
        <div class="collapse navbar-collapse" id="navbar">
            <table class="table table-condensed">
                <th class="text-center">Taille Disque Total</th>
                <th class="text-center">Taille utilisé</th>
                <th class="text-center">Taille Restante</th>
                <th class="text-center">% Utilisé</th>
                <th class="text-center">Disque</th>
                <tr>
                    {{ range $k,$v := .spacedisk}}
                    {{ if compare_not $k "0"}}
                    <td class="text-center">{{ $v }}</td>
                    {{ end }}
                    {{ end }}
                </tr>
            </table>
            <a class="navbar-brand" href="/">{{ .title }}</a>
            <ul class="nav navbar-nav">

                <li>
                    <a data-toggle="modal" data-target="#mkdir" href="#">
                        Créer un nouveau dossier
                    </a>
                </li>
                <li>
                    <a data-toggle="modal" data-target="#chown" href="#">
                        Changer les droits
                    </a>
                </li>

            </ul>
            <div class="navbar-nav navbar-form">

            <ol class="breadcrumb">
                <li><a href="/">Home</a></li>
                {{ range $k, $c := .chemin }}
                <li>
                    {{ $c }}
                </li>
                {{ end }}
            </ol>
            </div>
            <form action="/cmdperso{{ .href }}" method="post" class="navbar-form navbar-right" role="search">
                <div class="form-group">
                    <input type="input" class="form-control" placeholder="Cmd Perso : " size="50" name="cmdperso" value="" />
                </div>
                <button type="submit" class="btn btn-primary">
                    Valider
                </button>
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

<nav class="navbar navbar-default navbar-fixed-top" role="navigation">
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
                <button type="button" class="btn btn-danger" data-toggle="modal" data-target="#DtsToAc3">
                    Dts to AC3
                </button>
            </div>
            <div class="navbar-form navbar-left">
                <button type="button" class="btn btn-danger" data-toggle="modal" data-target="#KeepOneAudio">
                    Garder une seule piste audio
                </button>
            </div>
        </div><!-- /.navbar-collapse -->
    </div><!-- /.container-fluid -->
</nav>
<br>
<br>

<h2>{{ .name }}</h2>
<a href="/list/{{ .back }}"><span class="glyphicon glyphicon-arrow-up"></span></a>

<div class="col-md-12">
    <form method="POST">
        <fieldset class="form-group col-md-8">
            <label for="exampleSelect1">Changer de dossier</label>
            <select class="form-control" name="newPath" size="15">
                {{ range $l := .listFolder }}
                <option value="{{ $l }}">{{ $l }}</option>
                {{ end }}
            </select>
        </fieldset>

        <button type="submit" class="btn btn-primary col-md-5">Submit</button>
    </form>
</div>
<div class="col-md-12">
{{ .htmlMediaInfo | str2html }}
</div>


<div class="modal fade" id="KeepOneAudio" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
    <div class="modal-dialog" role="document">
        <form method="POST" class="form-horizontal col-sm-12" action="/oneaudio/{{ .back }}/{{ .name }}">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title" id="myModalLabel">Garder une seule Piste audio</h4>
                </div>
                <div class="modal-body">
                    <fieldset>
                        <div class="col-sm-12 form-group">
                            <label for="newName">Choisir un nouveau nom : </label>
                            <input type="text" class="form-control" id="newName" placeholder="Choisi un nouveau nom" name="newName" value="{{ .name }}" />
                        </div>
                    </fieldset>
                    <fieldset>
                        <div class="col-sm-12 form-group">
                            <label for="number">Choisir la piste à garder : Audio# ?</label>
                            <input type="number" class="form-control" id="number" placeholder="ex : Audio# : 1" name="number" value="" />
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
<div class="modal fade" id="DtsToAc3" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
    <div class="modal-dialog" role="document">
        <form method="POST" class="form-horizontal col-sm-12" action="/dtstoac3/{{ .back }}/{{ .name }}">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title" id="myModalLabel">Dts to AC3</h4>
                </div>
                <div class="modal-body">
                    <fieldset>
                        <div class="col-sm-12 form-group">
                            <label for="dtstoac3">Choisir un nouveau nom : </label>
                            <input type="text" class="form-control" id="dtstoac3" placeholder="Choisi un nouveau nom" name="dtstoac3" value="{{ .name }}" />
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

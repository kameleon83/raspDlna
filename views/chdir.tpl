
<h2>{{ .name }}</h2>
<a href="/list/{{ .back }}"><span class="glyphicon glyphicon-arrow-up"></span></a>

<div class="col-md-12">
    <form method="POST">
        <fieldset class="form-group">
            <label for="exampleSelect1">Changer de dossier</label>
            <select class="form-control" name="newPath">
                {{ range $l := .listFolder }}
                <option value="{{ $l }}">{{ $l }}</option>
                {{ end }}
            </select>
        </fieldset>
        <fieldset class="form-group">
            <label for="cheminComplet1">Chemin complet</label>
            <textarea class="form-control" id="cheminComplet1" rows="4" placeholder="Modifier Chemin si besoin" value="{{ .back}}/{{ .name }}" name="textarea-chemin">{{ .root }}{{ .back}}/{{ .name }}</textarea>
        </fieldset>

        <button type="submit" class="btn btn-primary">Submit</button>
    </form>
</div>
<br>
<br>
<br>
<br>
<div class="col-md-12">

    <a href="/srt/{{ .edit | str2html }}" class="btn btn-info" name="button">Créer Sous-titre</a>
</div>

<br>
<br>
<br>
<br>
<div class="">

{{ .htmlMediaInfo | str2html }}
</div>

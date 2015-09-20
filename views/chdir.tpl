
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

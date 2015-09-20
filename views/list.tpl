<h1>{{ .title }}</h1>

<nav>
	<ol class="breadcrumb">
		<li><a href="/">Home</a></li>
		{{ range $k, $c := .chemin }}
		<li>
			{{ $c }}
		</li>
		{{ end }}
	</ol>
</nav>

<div class="">

	<table class="table table-striped">
		<tr>
			<th>
				<a href="/list/{{ .back }}"><span class="glyphicon glyphicon-arrow-up"></span></a>
			</th>
			<th>Nom</th>
			<th>Taille</th>
			<th>Srt</th>
			<th>Edition</th>
		</tr>
		{{ range $k, $l := .dirname }}
		<tr>
			<td>
				{{ $a := substr $l.Name 0 1}}
				{{ if compare $a "." }}
				<span class="glyphicon glyphicon-ok"></span>
				{{ end }}
			</td>
			<td>
				{{ if $l.IsDir }}
				<a href='/list{{ $.href}}/{{ $l.Name }}'>

					{{ $l.Name }}
				</a>
				{{ else }}
				{{ $l.Name }}<span class="label label-default pull-right">{{ $l.NameExt }}</span>
				{{ end }}
			</td>
			<td>
				{{ $l.Size }} . {{ $l.NameSize }}
			</td>
			<td>
				{{ if compare $l.Srt "1"}}
				<span class="glyphicon glyphicon-list-alt"></span>.
				{{ $l.SizeSrt }} {{ $l.NameTailleSrt }}
				{{ end}}
			</td>
			<td class="inline-td">
				<form method="POST" action="/vues{{ $.href }}/{{ $l.Name}}">
					{{ if not $l.IsDir }}
					<button type="submit" class="btn btn-link ">
						<span class="glyphicon glyphicon-eye-open"></span>
					</button>
					{{ end }}
				</form>

				<form method="POST" action="/pas-vues{{ $.href }}/{{ $l.Name }}">
					{{ if not $l.IsDir }}
					<button type="submit" class="btn btn-link ">
						<span class="glyphicon glyphicon-eye-close"></span>
					</button>

					{{ end }}
				</form>
				{{ if not $l.IsDir }}
				<button type="button" class="btn btn-link">
					<a href="/srt{{ $.href }}/{{$l.Name }}">
						<span class="glyphicon glyphicon-list-alt"></span>
                    </a>
				</button>
				{{ else }}
				<button type="button" class="btn btn-link">
					<a href='/list{{ $.href}}/{{ $l.Name }}'>
						<span class="glyphicon glyphicon-folder-close" aria-hidden="true"></span>
					</a>
				</button>
				{{ end }}
				<button type="button" class="btn btn-link">
					<a href="/edit{{ $.href }}/{{$l.Name }}">
						<span class="glyphicon glyphicon-pencil"></span>
                    </a>
				</button>
				<button type="button" class="btn btn-link">
					<a href="/delete{{ $.href }}/{{$l.Name }}">
						<span class="glyphicon glyphicon-trash"></span>
					</a>
				</button>
				<button type="button" class="btn btn-link" data-toggle="modal" data-target="#myModal{{ $k }}">
					<span class="glyphicon glyphicon-tags"></span>
				</button>

				<div class="modal fade" id="myModal{{ $k }}" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
					<div class="modal-dialog" role="document">
						<form method="POST" class="form-horizontal col-sm-12" action="/rename{{ $.href }}/{{$l.Name }}">
							<div class="modal-content">
								<div class="modal-header">
									<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
									<h4 class="modal-title" id="myModalLabel">Renommer le {{ if $l.IsDir }}dossier{{ else }}fichier{{ end }} : <br/><br/>{{ $l.Name }}</h4>
								</div>
								<div class="modal-body">
									<div class="form-group">
										<fieldset class="form-group">
											<label for="rename" class="col-sm-2 control-label">Nom</label>
											<div class="col-sm-10">
												<input type="text" class="form-control" id="rename"  placeholder="Renommer" name="rename" value="{{ $l.Name }}" />
											</div>
										</fieldset>
									</div>
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
			</td>
		</tr>
		{{ end }}
	</table>
</div>

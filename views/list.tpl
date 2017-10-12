<div class="margin-top">

</div>

{{ template "navbar.tpl" . }}

<div class="col-md-12">

	<table class="table table-striped table-hover">
		<thead>
			<tr class="row">
				<th class="col-md-1">
					<a href="/list/{{ .back }}"><span class="glyphicon glyphicon-arrow-up"></span></a>
				</th>
				<th class="col-md-1">User : Group</th>
				<th class="col-md-5">Nom</th>
				<th class="col-md-1">Taille</th>
				<th class="col-md-1">Srt</th>
				<th class="col-md-3">Edition</th>
			</tr>
		</thead>
		<tbody>

			{{ range $k, $l := .dirname }}
			<tr class="row">
				<td class="col-md-1">
					{{ $a := substr $l.Name 0 1}}
					{{ if compare $a "." }}
					<span class="glyphicon glyphicon-ok"></span>
					{{ end }}
				</td>
				<td class="col-md-1">
					{{ range $k,$v := $.user }}
					{{ if compare $v $l.GetUid }}
					{{ $k }}
					{{ end }}
					{{ end }}
					:
					{{ range $k,$v := $.group }}
					{{ if compare $v $l.GetGid }}
					{{ $k }}
					{{ end }}
					{{ end }}

				</td>
				<td class="col-md-5">
					{{ if $l.IsDir }}
					<a href='/list{{ $.href}}/{{ $l.Name }}'>

						{{ $l.Name }}
					</a>
					{{ else }}
					{{ $l.Name }}<span class="label label-default pull-right">{{ $l.NameExt }}</span>
					{{ end }}
				</td>
				<td class="col-md-1">
					{{ $l.Size }} . {{ $l.NameSize }}
				</td>
				<td class="col-md-1">
					{{ if compare $l.Srt "1"}}
					<span class="glyphicon glyphicon-list-alt"></span>.
					{{ $l.SizeSrt }} {{ $l.NameTailleSrt }}
					{{ end}}
				</td>
				<td class="inline-td col-md-3">
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
		</tbody>
	</table>
</div>

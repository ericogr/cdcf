package docker

import (
  "encoding/base64"
)

func CreateDockerConfigEncoded(usuarioHarbor string, senhaHarbor string, address string) (encodedDockerConfig string) {
  return encoderBase64(CreateDockerConfig(usuarioHarbor, senhaHarbor, address))
}

func CreateDockerConfig(usuarioHarbor string, senhaHarbor string, address string) (encodedDockerConfig string) {
  return templateDockerConfig(encoderBase64(concatenaUsuarioSenha(usuarioHarbor, senhaHarbor)), address)
}

func templateDockerConfig(auth string, address string) (dockerConfig string){
  dockerConfig = "{\"auths\":{\"" + address + "\":{\"auth\":\"" + auth + "\"}}}"
  return
}

func concatenaUsuarioSenha(usuario string, senha string) (concatenado string) {
  concatenado = usuario + ":" + senha
  return
}

func encoderBase64(texto string) (codificado string) {
  data := []byte(texto)
  codificado = base64.StdEncoding.EncodeToString(data)
  return
}

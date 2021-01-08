package iui

import (
	"fmt"

	"github.com/vktec/gll"
	"github.com/vktec/gll/glh"
)

const vertSrc = `
#version 150
void main() {
	int x = (gl_VertexID & 1) << 2;
	int y = (gl_VertexID & 2) << 1;
	gl_Position = vec4(x-1, y-1, 0, 1);
}
`

var shaderSources []string

func RegisterShader(fragSrc string) (id int) {
	shaderSources = append(shaderSources, fragSrc)
	return len(shaderSources) - 1
}

func buildShaders(gl gll.GL300) (ShaderRegistry, error) {
	reg := make(ShaderRegistry, len(shaderSources))
	vshad, err := glh.NewShader(gl, gll.VERTEX_SHADER, vertSrc)
	if err != nil {
		return nil, fmt.Errorf("Error building vertex shader: %w", err)
	}
	defer gl.DeleteShader(vshad)

	// FIXME: programs are not cleaned up on error
	for i, fragSrc := range shaderSources {
		fshad, err := glh.NewShader(gl, gll.FRAGMENT_SHADER, fragSrc)
		if err != nil {
			return nil, fmt.Errorf("Error building fragment shader: %w", err)
		}

		reg[i], err = glh.NewProgram(gl, vshad, fshad)
		gl.DeleteShader(fshad)
		if err != nil {
			return nil, fmt.Errorf("Error linking program: %w", err)
		}
	}
	return reg, nil
}

type ShaderRegistry []uint32

func (reg ShaderRegistry) GetShader(id int) uint32 {
	return reg[id]
}

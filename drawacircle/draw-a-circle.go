package main

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"strings"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	fps    = 10
	width  = 500
	height = 500

	vertexShaderSource = `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1, 1, 1, 1.0);
		}
	` + "\x00"
)

var (
	triangle = make([]float32, 9)

	// triangle = []float32{
	// 	// // Q1:
	// 	// 0, 200, 0,
	// 	// 50, 194, 0,
	// 	// 0, 0, 0,

	// 	// 50, 194, 0,
	// 	// 100, 173, 0,
	// 	// 0, 0, 0,

	// 	// 100, 173, 0,
	// 	// 150, 132, 0,
	// 	// 0, 0, 0,

	// 	// 150, 132, 0,
	// 	// 200, 0, 0,
	// 	// 0, 0, 0,

	// 	// // Q2:
	// 	// 0, 200, 0,
	// 	// -50, 194, 0,
	// 	// 0, 0, 0,

	// 	// -50, 194, 0,
	// 	// -100, 173, 0,
	// 	// 0, 0, 0,

	// 	// -100, 173, 0,
	// 	// -150, 132, 0,
	// 	// 0, 0, 0,

	// 	// -150, 132, 0,
	// 	// -200, 0, 0,
	// 	// 0, 0, 0,

	// 	// // Q3
	// 	// 0, -200, 0,
	// 	// -50, -194, 0,
	// 	// 0, 0, 0,

	// 	// -50, -194, 0,
	// 	// -100, -173, 0,
	// 	// 0, 0, 0,

	// 	// -100, -173, 0,
	// 	// -150, -132, 0,
	// 	// 0, 0, 0,

	// 	// -150, -132, 0,
	// 	// -200, 0, 0,
	// 	// 0, 0, 0,

	// 	// // Q4
	// 	// 0, -200, 0,
	// 	// 50, -194, 0,
	// 	// 0, 0, 0,

	// 	// 50, -194, 0,
	// 	// 100, -173, 0,
	// 	// 0, 0, 0,

	// 	// 100, -173, 0,
	// 	// 150, -132, 0,
	// 	// 0, 0, 0,

	// 	// 150, -132, 0,
	// 	// 200, 0, 0,
	// 	// 0, 0, 0,
	// }

	sections = 1
)

func generateTriangle(currentIndex float32, radius float32, total float32) []float32 {
	arr := make([]float32, 9)

	// find X
	x := currentIndex * (radius / total)

	// x2 + y2 = r2
	y := math.Sqrt((math.Pow(float64(radius), float64(2)) - math.Pow(float64(x), float64(2))))

	// find next point
	a := (currentIndex + 1) * (radius / total)

	// find y of the second point
	b := math.Sqrt((math.Pow(float64(radius), float64(2)) - math.Pow(float64(a), float64(2))))

	arr[0] = x
	arr[1] = float32(y)
	arr[2] = 0

	arr[3] = a
	arr[4] = float32(b)
	arr[5] = 0

	arr[6] = 0
	arr[7] = 0
	arr[8] = 0

	return arr
}

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	increase := true

	for !window.ShouldClose() {
		t := time.Now()

		if sections == 64 {
			increase = false
		}

		create()
		vao := makeVao(triangle)
		draw(vao, window, program)

		if sections == 1 {
			increase = true
		}

		if increase == true {
			sections++
		} else {
			sections--
		}

		time.Sleep(time.Second/time.Duration(2) - time.Since(t))
	}
}

func create() {
	for h := float32(0); h < float32(sections); h++ {
		quad1 := generateTriangle(h, 200, float32(sections))

		tmp1 := make([]float32, len(quad1))
		copy(tmp1, quad1)
		quad2 := generateQuad(tmp1, 2)

		tmp2 := make([]float32, len(quad1))
		copy(tmp2, quad1)
		quad3 := generateQuad(tmp2, 3)

		tmp3 := make([]float32, len(quad1))
		copy(tmp3, quad1)
		quad4 := generateQuad(tmp3, 4)

		triangle = append(triangle, quad1...)
		triangle = append(triangle, quad2...)
		triangle = append(triangle, quad3...)
		triangle = append(triangle, quad4...)
	}

	for i := 0; i < len(triangle); i++ {
		triangle[i] = (triangle[i] / 250)
	}
}

func generateQuad(triangle []float32, quad int32) []float32 {
	switch quad {
	case 2:
		triangle[0] = triangle[0] * -1
		triangle[3] = triangle[3] * -1

		return triangle

	case 3:
		triangle[0] = triangle[0] * -1
		triangle[1] = triangle[1] * -1
		triangle[3] = triangle[3] * -1
		triangle[4] = triangle[4] * -1

		return triangle

	case 4:
		triangle[1] = triangle[1] * -1
		triangle[4] = triangle[4] * -1

		return triangle
	}

	return triangle
}

func draw(vao uint32, window *glfw.Window, program uint32) {
	fmt.Println("I am here", sections)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

	glfw.PollEvents()
	window.SwapBuffers()
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Draw Points", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

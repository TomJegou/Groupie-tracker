import * as THREE from 'three';
import { GLTFLoader } from 'three/addons/loaders/GLTFLoader.js';
const canvas = document.querySelector('.webgl')
const scene = new THREE.Scene ()
const loader = new GLTFLoader ( )
loader.load('http://localhost/static/img/casque.gltf',function(gltf){
    console.log(gltf)
    const root = gltf.scene;
    root.scale.set(0.05,0.05,0.05)
    scene.add(root);
},function(error){
    console.log("An error occurred :",error)
},)

const light = new THREE.DirectionalLight(0xffffff,1)
light.position.set(2,2,5)
scene.add(light)

const sizes ={
    width: window.innerWidth,
    height: window.innerHeight
}

const camera = new THREE.PerspectiveCamera(75, sizes.width/sizes.height, 0.1,100)
camera.position.set(0,1,2)
scene.add(camera)

const renderer = new THREE.WebGL1Renderer({
    canvas: canvas
})
    camera.position.set(-50, 40, 350);
    renderer = new THREE.WebGLRenderer({antialias:true, alpha: true});
    renderer.setSize(container.clientWidth,container.clientHeight);
    renderer.setPixelRatio(window.devicePixelRatio)
    container.appendChild(renderer.domElement);

renderer.setSize(size.width,size.height)
renderer.setPixelRatio(Math.min(window.devicePixelRatio,2))
renderer.shadowMap.enable = true
renderer.gammaOuput = true

renderer.render(scene.camera)
function animate(){
}


animate()
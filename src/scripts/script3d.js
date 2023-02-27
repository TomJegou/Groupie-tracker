import{
    three
}from 'three';

let container;
let camera;
let scene;
let house;

function init(){
    container = document.querySelector('.scene')
    scene = new THREE.Scene('.scene');
    const fov = 35;
    const aspect = container.clientWidth / container.clientHeight;
    const near = 0.1;
    const far = 500;
    camera = new THREE.PerspectiveCamera(fov,aspect,near,far)
}
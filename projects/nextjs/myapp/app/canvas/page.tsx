'use client';

import React, { useEffect } from 'react';
import './canvas.css';
export default function CanvasPage() {
    const speed = 0.01
    useEffect(() => {
        const animateElement = document.getElementById('animate1');
        let angle = 0;

        const updatePosition = () => {
            if (animateElement) {
                // Define the radius and speed of the circular motion
                const radius = 30; // in pixels

                // Calculate the new position
                const x = Math.cos(angle) * radius;
                const y = Math.cos(angle) * radius;

                // Update the position of the element
                animateElement.style.transform = `translate(${x}px, ${y}px)`;

                // Increment the angle for the next frame
                angle += speed;

                // Schedule the next frame
                requestAnimationFrame(updatePosition);
            }
        };

        // Start the animation
        updatePosition();
    }, []);
    // useEffect(() => {
    //     const animateElement = document.getElementById('animate2');
    //     let angle = 0;

    //     const updatePosition = () => {
    //         if (animateElement) {
    //             // Define the radius and speed of the circular motion
    //             const radius = 30; // in pixels

    //             // Calculate the new position
    //             const x = - Math.cos(angle) * radius;
    //             const y = Math.cos(angle) * radius;

    //             // Update the position of the element
    //             animateElement.style.transform = `translate(${x}px, ${y}px)`;

    //             // Increment the angle for the next frame
    //             angle += speed;

    //             // Schedule the next frame
    //             requestAnimationFrame(updatePosition);
    //         }
    //     };

    //     // Start the animation
    //     updatePosition();
    // }, []);
    // useEffect(() => {
    //     const animateElement = document.getElementById('animate3');
    //     let angle = 0;

    //     const updatePosition = () => {
    //         if (animateElement) {
    //             // Define the radius and speed of the circular motion
    //             const radius = 30; // in pixels

    //             // Calculate the new position
    //             const x = Math.cos(angle) * radius;
    //             const y = - Math.cos(angle) * radius;

    //             // Update the position of the element
    //             animateElement.style.transform = `translate(${x}px, ${y}px)`;

    //             // Increment the angle for the next frame
    //             angle += speed;

    //             // Schedule the next frame
    //             requestAnimationFrame(updatePosition);
    //         }
    //     };

    //     // Start the animation
    //     updatePosition();
    // }, []);
    // useEffect(() => {
    //     const animateElement = document.getElementById('animate4');
    //     let angle = 0;

    //     const updatePosition = () => {
    //         if (animateElement) {
    //             // Define the radius and speed of the circular motion
    //             const radius = 30; // in pixels

    //             // Calculate the new position
    //             const x = - Math.cos(angle) * radius;
    //             const y = - Math.cos(angle) * radius;

    //             // Update the position of the element
    //             animateElement.style.transform = `translate(${x}px, ${y}px)`;

    //             // Increment the angle for the next frame
    //             angle += speed;

    //             // Schedule the next frame
    //             requestAnimationFrame(updatePosition);
    //         }
    //     };

    //     // Start the animation
    //     updatePosition();
    // }, []);
    return (
        <div className="canvas">
            <div className="animate" id="animate1"></div>
            <div className="animate" id="animate2"></div>
            <div className="animate" id="animate3"></div>
            <div className="animate" id="animate4"></div>
        </div>
    );
}
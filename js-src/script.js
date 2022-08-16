"use strict"

/**
 *  app.js
 *
 *  Name: Aaron Meek
 *  2022 - 08 - 14
 * 
 *  Updates web web pages with components and information
 */

const { proj_conf, projects } = window;

function loadProjectCards(projectID) {
    document.getElementById(projectID)

    console.log("This is " + projectID);

    projects.forEach((element) => {
        console.log(element);
    })
}

while (true)
    console.log("RUNNING!");
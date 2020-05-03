#!/bin/bash
go build ./cmd/ray-tracer/
./ray-tracer --width 1920 --height 1080 --number_of_samples 10 --output_file_name "./example_images/simple_scene_1920_1080_10.png" --scene "simple_scene" --show_after_complete 0 --number_of_workers 4 > ./log/simple_scene_1920_1080_10.log
./ray-tracer --width 1920 --height 1080 --number_of_samples 25 --output_file_name "./example_images/simple_scene_1920_1080_25.png" --scene "simple_scene" --show_after_complete 0 --number_of_workers 4 > ./log/simple_scene_1920_1080_25.log
./ray-tracer --width 1920 --height 1080 --number_of_samples 100 --output_file_name "./example_images/simple_scene_1920_1080_100.png" --scene "simple_scene" --show_after_complete 0 --number_of_workers 4 > ./log/simple_scene_1920_1080_100.log

./ray-tracer --width 555 --height 555 --number_of_samples 10 --output_file_name "./example_images/cornell_box_555_555_10.png" --scene "cornell_box" --show_after_complete 0 --number_of_workers 4 > ./log/cornell_box_555_555_10.log
./ray-tracer --width 555 --height 555 --number_of_samples 25 --output_file_name "./example_images/cornell_box_555_555_25.png" --scene "cornell_box" --show_after_complete 0 --number_of_workers 4 > ./log/cornell_box_555_555_25.log
./ray-tracer --width 555 --height 555 --number_of_samples 100 --output_file_name "./example_images/cornell_box_555_555_100.png" --scene "cornell_box" --show_after_complete 0 --number_of_workers 4 > ./log/cornell_box_555_555_100.log
./ray-tracer --width 555 --height 555 --number_of_samples 1000 --output_file_name "./example_images/cornell_box_555_555_1000.png" --scene "cornell_box" --show_after_complete 0 --number_of_workers 4 > ./log/cornell_box_555_555_1000.log

./ray-tracer --width 600 --height 600 --number_of_samples 10 --output_file_name "./example_images/cornell_box_octahedron_600_600_10.png" --scene "cornell_box_octahedron" --show_after_complete 0 --number_of_workers 4 > ./log/cornell_box_octahedron_600_600_10.log
./ray-tracer --width 600 --height 600 --number_of_samples 25 --output_file_name "./example_images/cornell_box_octahedron_600_600_25.png" --scene "cornell_box_octahedron" --show_after_complete 0 --number_of_workers 4 > ./log/cornell_box_octahedron_600_600_25.log
./ray-tracer --width 600 --height 600 --number_of_samples 100 --output_file_name "./example_images/cornell_box_octahedron_600_600_100.png" --scene "cornell_box_octahedron" --show_after_complete 0 --number_of_workers 4 > ./log/cornell_box_octahedron_600_600_100.log
./ray-tracer --width 600 --height 600 --number_of_samples 1000 --output_file_name "./example_images/cornell_box_octahedron_600_600_1000.png" --scene "cornell_box_octahedron" --show_after_complete 0 --number_of_workers 4 > ./log/cornell_box_octahedron_600_600_1000.log


./ray-tracer --width 1920 --height 1080 --number_of_samples 10 --output_file_name "./example_images/simple_scene_spheres_1920_1080_10.png" --scene "simple_scene_spheres" --show_after_complete 0 --number_of_workers 4 > ./log/simple_scene_spheres_1920_1080_10.log
./ray-tracer --width 1920 --height 1080 --number_of_samples 25 --output_file_name "./example_images/simple_scene_spheres_1920_1080_25.png" --scene "simple_scene_spheres" --show_after_complete 0 --number_of_workers 4 > ./log/simple_scene_spheres_1920_1080_25.log
./ray-tracer --width 1920 --height 1080 --number_of_samples 100 --output_file_name "./example_images/simple_scene_spheres_1920_1080_100.png" --scene "simple_scene_spheres" --show_after_complete 0 --number_of_workers 4 > ./log/simple_scene_spheres_1920_1080_100.log
./ray-tracer --width 1920 --height 1080 --number_of_samples 1000 --output_file_name "./example_images/simple_scene_spheres_1920_1080_1000.png" --scene "simple_scene_spheres" --show_after_complete 0 --number_of_workers 4 > ./log/simple_scene_spheres_1920_1080_1000.log






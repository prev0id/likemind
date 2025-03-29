-- +goose Up
-- +goose StatementBegin
INSERT INTO interest_groups (id, name) VALUES
  (1, 'Arts & Crafts'),
  (2, 'Music & Performing Arts'),
  (3, 'Sports & Fitness'),
  (4, 'Outdoor & Nature'),
  (5, 'Technology & Gaming'),
  (6, 'Travel & Adventure'),
  (7, 'Food & Culinary'),
  (8, 'Reading & Writing'),
  (9, 'Social & Community'),
  (10, 'Mind & Strategy'),
  (11, 'Miscellaneous');

INSERT INTO interests (id, name, description, group_id) VALUES
  (1, 'Painting', 'Create visual art using various paints.', 1),
  (2, 'Drawing', 'Express creativity with pencils, pens, or charcoal.', 1),
  (3, 'Sculpting', 'Mold or carve materials into three-dimensional art.', 1),
  (4, 'Calligraphy', 'Practice artistic, stylized handwriting.', 1),
  (5, 'Photography', 'Capture moments and scenery with a camera.', 1),
  (6, 'Scrapbooking', 'Assemble photos and mementos in creative albums.', 1),
  (7, 'Pottery', 'Shape and fire clay to create ceramics.', 1),
  (8, 'Knitting', 'Craft garments or decor using yarn and needles.', 1),
  (9, 'Sewing', 'Stitch fabrics to create clothing or accessories.', 1),
  (10, 'Embroidery', 'Decorate fabric with needle and thread designs.', 1),
  (11, 'Jewelry Making', 'Design and construct wearable art pieces.', 1),
  (12, 'Origami', 'Fold paper into decorative or intricate designs.', 1),
  (13, 'Digital Art', 'Create art using digital tools and software.', 1),
  (14, 'Collage Art', 'Assemble diverse materials to form new artwork.', 1),
  (15, 'Woodworking', 'Craft items from wood with tools and creativity.', 1),
  (16, 'Candle Making', 'Create custom, scented candles.', 1),
  (17, 'Soap Making', 'Craft unique handmade soaps.', 1),
  (18, 'Glassblowing', 'Shape molten glass into artistic forms.', 1),
  (19, 'Mosaic Art', 'Piece together small fragments to form images.', 1),
  (20, 'Weaving', 'Interlace threads to form fabric or decorative items.', 1);

INSERT INTO interest (id, name, description, group_id) VALUES
  (21, 'Playing Guitar', 'Strum acoustic or electric guitars.', 2),
  (22, 'Playing Piano', 'Perform and compose on a keyboard instrument.', 2),
  (23, 'Singing', 'Use your voice for performance and expression.', 2),
  (24, 'Playing Drums', 'Keep rhythm with percussion instruments.', 2),
  (25, 'Violin Playing', 'Create melodies on the violin.', 2),
  (26, 'DJing', 'Mix and play tracks at events or online.', 2),
  (27, 'Dancing', 'Express yourself through movement and rhythm.', 2),
  (28, 'Acting', 'Perform in theater, film, or television.', 2),
  (29, 'Stand-up Comedy', 'Deliver humorous monologues to an audience.', 2),
  (30, 'Magic Tricks', 'Perform illusions and sleight-of-hand magic.', 2),
  (31, 'Musical Composition', 'Write original music and scores.', 2),
  (32, 'Beatboxing', 'Create rhythmic sounds using only your voice.', 2),
  (33, 'Karaoke', 'Sing along to popular songs in a fun setting.', 2),
  (34, 'Playing Bass', 'Support music with deep, rhythmic bass lines.', 2),
  (35, 'Flute Playing', 'Produce light, melodic tunes with a flute.', 2);

INSERT INTO interest (id, name, description, group_id) VALUES
  (36, 'Running', 'Improve fitness with jogging or sprints.', 3),
  (37, 'Cycling', 'Enjoy road or mountain biking for exercise.', 3),
  (38, 'Swimming', 'Engage in a full-body water workout.', 3),
  (39, 'Yoga', 'Enhance flexibility and mindfulness.', 3),
  (40, 'Weightlifting', 'Build strength through resistance training.', 3),
  (41, 'Hiking', 'Explore nature on foot along trails.', 3),
  (42, 'Rock Climbing', 'Scale natural or artificial rock walls.', 3),
  (43, 'Martial Arts', 'Practice combat techniques and self-defense.', 3),
  (44, 'Dancing (Fitness)', 'Enjoy dance-based workout routines.', 3),
  (45, 'Rowing', 'Exercise using oars in watercraft or machines.', 3),
  (46, 'Boxing', 'Develop fitness and technique in combat sport.', 3),
  (47, 'Tennis', 'Play an engaging racket sport.', 3),
  (48, 'Basketball', 'Team up for fast-paced court play.', 3),
  (49, 'Soccer', 'Enjoy team strategy on the field.', 3),
  (50, 'Golfing', 'Practice precision on the golf course.', 3),
  (51, 'Skiing', 'Glide over snowy slopes in winter.', 3),
  (52, 'Snowboarding', 'Carve paths down snow-covered hills.', 3),
  (53, 'Surfing', 'Ride ocean waves on a surfboard.', 3),
  (54, 'Skateboarding', 'Perform tricks on a skateboard.', 3),
  (55, 'Pilates', 'Strengthen core muscles with low-impact exercise.', 3),
  (56, 'Crossfit', 'Engage in high-intensity interval training.', 3),
  (57, 'Badminton', 'Enjoy a fast-paced indoor racket sport.', 3),
  (58, 'Table Tennis', 'Play competitive, quick-paced ping pong.', 3),
  (59, 'Martial Arts (Advanced)', 'Explore advanced combat techniques.', 3),
  (60, 'Fencing', 'Compete in swordplay with agility and precision.', 3);

INSERT INTO interest (id, name, description, group_id) VALUES
  (61, 'Gardening', 'Cultivate plants and design landscapes.', 4),
  (62, 'Bird Watching', 'Observe and identify various bird species.', 4),
  (63, 'Fishing', 'Enjoy leisure or sport fishing in different settings.', 4),
  (64, 'Camping', 'Experience outdoor living with overnight stays.', 4),
  (65, 'Astronomy', 'Study stars, planets, and celestial events.', 4),
  (66, 'Geocaching', 'Hunt for hidden treasures using GPS.', 4),
  (67, 'Hiking (Extended)', 'Explore more challenging trails.', 4),
  (68, 'Rock Collecting', 'Gather and study unique rock specimens.', 4),
  (69, 'Nature Photography', 'Capture the beauty of the natural world.', 4),
  (70, 'Kayaking', 'Paddle through rivers, lakes, or coastal waters.', 4),
  (71, 'Canoeing', 'Navigate waterways in a small, narrow boat.', 4),
  (72, 'Sailing', 'Learn to navigate and operate a sailboat.', 4),
  (73, 'Horseback Riding', 'Enjoy riding and caring for horses.', 4),
  (74, 'Foraging', 'Collect wild edibles in nature.', 4),
  (75, 'Outdoor Fitness', 'Exercise using natural surroundings.', 4),
  (76, 'Trail Running', 'Run on off-road natural paths.', 4);

INSERT INTO interest (id, name, description, group_id) VALUES
  (77, 'Video Gaming', 'Play digital games across platforms.', 5),
  (78, 'Board Gaming', 'Enjoy strategic and cooperative tabletop games.', 5),
  (79, 'Puzzle Solving', 'Challenge your mind with puzzles.', 5),
  (80, 'Coding', 'Develop software and learn programming languages.', 5),
  (81, 'Robotics', 'Build and program robotic systems.', 5),
  (82, 'Virtual Reality Gaming', 'Experience immersive digital worlds.', 5),
  (83, 'App Development', 'Create mobile or web applications.', 5),
  (84, '3D Printing', 'Design and print three-dimensional objects.', 5),
  (85, 'Streaming', 'Broadcast gaming or creative content live.', 5),
  (86, 'Drone Flying', 'Operate and pilot unmanned aerial vehicles.', 5),
  (87, 'Computer Building', 'Assemble and customize personal computers.', 5),
  (88, 'eSports', 'Compete in professional video game tournaments.', 5),
  (89, 'Puzzle Video Games', 'Enjoy games focused on mental challenges.', 5),
  (90, 'Retro Gaming', 'Play classic games from past decades.', 5);

INSERT INTO interest (id, name, description, group_id) VALUES
  (91, 'Backpacking', 'Travel light on extended journeys.', 6),
  (92, 'Road Trips', 'Explore new places by car.', 6),
  (93, 'Cultural Tourism', 'Immerse in the traditions of different cultures.', 6),
  (94, 'Camping (Travel)', 'Combine travel with outdoor camping experiences.', 6),
  (95, 'Sightseeing', 'Visit and admire notable landmarks.', 6),
  (96, 'Travel Blogging', 'Document and share travel adventures online.', 6),
  (97, 'Adventure Sports', 'Engage in adrenaline-inducing activities.', 6),
  (98, 'Scuba Diving', 'Explore underwater ecosystems.', 6),
  (99, 'Paragliding', 'Glide through the air using a parachute-like canopy.', 6),
  (100, 'Hot Air Ballooning', 'Float peacefully in a colorful balloon.', 6),
  (101, 'Train Travel', 'Enjoy scenic journeys by rail.', 6),
  (102, 'Wildlife Safaris', 'Observe animals in their natural habitats.', 6),
  (103, 'Culinary Travel', 'Discover global cuisines while traveling.', 6),
  (104, 'City Exploration', 'Experience the charm of urban life.', 6),
  (105, 'Nature Trekking', 'Hike through rugged and wild terrains.', 6);

INSERT INTO interest (id, name, description, group_id) VALUES
  (106, 'Cooking', 'Experiment with recipes and ingredients.', 7),
  (107, 'Baking', 'Create breads, pastries, and desserts.', 7),
  (108, 'Grilling', 'Enjoy outdoor cooking on a grill.', 7),
  (109, 'Wine Tasting', 'Sample and evaluate different wines.', 7),
  (110, 'Craft Beer Brewing', 'Brew personalized batches of beer.', 7),
  (111, 'Mixology', 'Craft and experiment with cocktails.', 7),
  (112, 'Food Blogging', 'Share culinary adventures and recipes online.', 7),
  (113, 'Vegan Cooking', 'Explore plant-based recipes and dishes.', 7),
  (114, 'Fermenting Foods', 'Make homemade pickles, kimchi, and more.', 7),
  (115, 'Chocolate Making', 'Create artisan chocolate treats.', 7),
  (116, 'Cake Decorating', 'Design visually appealing cakes.', 7),
  (117, 'International Cuisine', 'Explore cooking styles from around the world.', 7),
  (118, 'Home Brewing', 'Brew beer or other beverages at home.', 7);

INSERT INTO interest (id, name, description, group_id) VALUES
  (119, 'Reading', 'Enjoy literature, fiction, and non-fiction works.', 8),
  (120, 'Creative Writing', 'Craft stories, poems, or essays.', 8),
  (121, 'Blogging', 'Write online posts on various topics.', 8),
  (122, 'Journaling', 'Record thoughts and daily experiences.', 8),
  (123, 'Poetry Writing', 'Compose short or long poetic works.', 8),
  (124, 'Book Clubbing', 'Discuss and analyze books with others.', 8),
  (125, 'Scriptwriting', 'Write for film, TV, or theater.', 8),
  (126, 'Essay Writing', 'Develop persuasive or descriptive essays.', 8),
  (127, 'Novella Writing', 'Create shorter forms of novel-length stories.', 8),
  (128, 'Short Story Writing', 'Craft concise fictional narratives.', 8),
  (129, 'Comic Writing', 'Develop stories in a graphic format.', 8);

INSERT INTO interest (id, name, description, group_id) VALUES
  (130, 'Volunteering', 'Donate time to community projects.', 9),
  (131, 'Mentoring', 'Guide and support others in personal growth.', 9),
  (132, 'Networking', 'Build professional and personal connections.', 9),
  (133, 'Social Dancing', 'Enjoy partner dancing in social settings.', 9),
  (134, 'Community Theater', 'Participate in local stage productions.', 9),
  (135, 'Event Planning', 'Organize gatherings or special events.', 9),
  (136, 'Blogging (Social)', 'Share experiences and ideas with a community.', 9),
  (137, 'Public Speaking', 'Engage and inform an audience.', 9),
  (138, 'Meetup Organizing', 'Create local groups for shared interests.', 9);

INSERT INTO interest (id, name, description, group_id) VALUES
  (139, 'Chess', 'Engage in a classic game of strategy.', 10),
  (140, 'Sudoku', 'Solve number-based puzzles for mental exercise.', 10),
  (141, 'Crosswords', 'Challenge yourself with word puzzles.', 10),
  (142, 'Brain Teasers', 'Enjoy puzzles that test your logic.', 10),
  (143, 'Meditation', 'Practice mindfulness and inner calm.', 10),
  (144, 'Debate Clubs', 'Hone your argumentative and public speaking skills.', 10),
  (145, 'Strategy Games', 'Play games that require careful planning.', 10),
  (146, 'Card Games', 'Enjoy various card games for fun and challenge.', 10);

INSERT INTO interest (id, name, description, group_id) VALUES
  (147, 'Pet Care', 'Spend time nurturing and caring for pets.', 11),
  (148, 'DIY Home Improvement', 'Tackle small renovation or craft projects.', 11),
  (149, 'Collecting', 'Gather stamps, coins, or other unique items.', 11),
  (150, 'Scrap Metal Art', 'Create art from recycled metal pieces.', 11),
  (151, 'Genealogy', 'Research and document family history.', 11),
  (152, 'Magic & Illusions', 'Learn and perform sleight-of-hand tricks.', 11),
  (153, 'Aromatherapy', 'Use essential oils for relaxation and wellness.', 11),
  (154, 'Homebrewing (Advanced)', 'Experiment with brewing beverages at a deeper level.', 11);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
select 1;
-- +goose StatementEnd

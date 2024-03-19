﻿using Explorer.Blog.API.Dtos;
using Explorer.Blog.API.Public;
using Explorer.BuildingBlocks.Core.Domain;
using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.Public.Administration;
using Explorer.Tours.Core.UseCases.Administration;
using FluentResults;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;
using System.Net.Http;

namespace Explorer.API.Controllers.Tourist
{
    [Authorize(Policy = "touristPolicy")]
    [Route("api/tourist/blog")]
    public class BlogController : BaseApiController
    {
        private readonly IBlogService _blogService;
        //private readonly ICommentService _commentService;

        public BlogController(IBlogService blogService)
        {
            _blogService = blogService;
            //_commentService = commentService;
        }

        /*[HttpPost]
        public ActionResult<BlogDto> Create([FromBody] BlogDto blog)
        {
            var result = _blogService.Create(blog);
            return CreateResponse(result);
        }*/


        [HttpPost]
        public async Task<ActionResult<BlogDto>> CreateAsync([FromBody] BlogDto blogDto)
        {
            try
            {             
                var blog = await _blogService.CreateBlogAsync(blogDto);
                return CreateResponse(blog);

            }
            catch (Exception e)
            {
                return StatusCode(500, e.Message);
            }
        }

        /*[HttpGet]
        public ActionResult<List<BlogDto>> GetAll()
        {
            var result = _blogService.GetAll();
            return CreateResponse(result);
        }*/

        [HttpGet]
        public async Task<ActionResult<List<BlogDto>>> GetAllAsync()
        {
            try
            {
                var blogDto = await _blogService.GetAllBlogsAsync();
                return CreateResponse(blogDto);
            }
            catch (Exception e)
            {
                return StatusCode(500, e.Message);
            }
        }

        /*[HttpGet("{id:int}")]
        public ActionResult<BlogDto> Get(int id)
        {
            var result = _blogService.Get(id);
            return CreateResponse(result);
        }*/

        [HttpGet("{id:int}")]
        public async Task<ActionResult<BlogDto>> GetAsync(int id)
        {
            try
            {
                var blogDto =  await _blogService.GetBlogByIdAsync(id);
                if (blogDto == null) return NotFound();
                return CreateResponse(blogDto);
            }
            catch (Exception e)
            {
                return StatusCode(500, e.Message);
            }
        }


        [HttpPut]
        public ActionResult<List<BlogDto>> UpdateBlog([FromBody] BlogDto blog)
        {
            var result = _blogService.Update(blog);
            var returnresult = _blogService.GetAll();

            return CreateResponse(returnresult);
        }

        /*[HttpPut("oneBlogUpdated")]
        public ActionResult<BlogDto> UpdateOneBlog([FromBody] BlogDto blog)
        {
            var result = _blogService.Update(blog);

            return CreateResponse(result);
        }*/

        [HttpPut("oneBlogUpdated")]
        public async Task<ActionResult<BlogDto>> UpdateOneBlog([FromBody] BlogDto blogDto)
        {
            try
            {
                var blog = await _blogService.UpdateBlogAsync(blogDto);
                return CreateResponse(blog);

            }
            catch (Exception e)
            {
                return StatusCode(500, e.Message);
            }
        }

        [HttpPost("createComment")]
        public ActionResult<CommentDto> Create([FromBody] CommentDto commentDto)
        {
            var result = _blogService.CreateComment(commentDto);
            return CreateResponse(result);
        }

        [HttpGet("comment/{id:int}")]
        public ActionResult<CommentDto> GetComment(int id)
        {
            var result = _blogService.GetComment(id);
            return CreateResponse(result);
        }

        [HttpPut("editComment")]
        public ActionResult<CommentDto> UpdateComment([FromBody] CommentDto commentDto)
        {
            var result = _blogService.UpdateComment(commentDto);
            return CreateResponse(result);
        }

        [HttpDelete("deleteComment/{id:int}")]
        public ActionResult DeleteComment(int id)
        {
            var result = _blogService.DeleteComment(id);
            return CreateResponse(result);
        }

        [HttpGet("allComments")]
        public ActionResult<PagedResult<CommentDto>> GetAllComments([FromQuery] int page, [FromQuery] int pageSize)
        {
            var result = _blogService.GetPagedComments(page, pageSize);
            return CreateResponse(result);
        }

        [HttpGet("blogComments/{blogId:int}")]
        public ActionResult<List<CommentDto>> GetCommentsByBlogId(int blogId)
        {
            var result = _blogService.GetCommentsByBlogId(blogId);
            return CreateResponse(result);
        }
        
        [HttpDelete("rating/{userId:int}/{blogId:int}")]
        public ActionResult DeleteRating(int blogId,int userId)
        {
            var result=_blogService.DeleteRating(blogId, userId);
            return CreateResponse(result);
        }

        [HttpPut("rating/{userId:int}/{blogId:int}/{value:int}")]
        public ActionResult<BlogDto> UpdateRating(int blogId, int userId,int value)
        {
            var result = _blogService.UpdateRating(blogId, userId,value);
            return CreateResponse(result);
        }

        /*[HttpGet("getByStatus/{state:int}")]
        public ActionResult<List<BlogDto>> GetBlogsByStatus(int state)
        {
            var result = _blogService.GetBlogsByStatus(state);
            return CreateResponse(result);
        }*/
        [HttpGet("getByStatus/{state:int}")]
        public async Task<ActionResult<List<BlogDto>>> GetBlogsByStatus(int state)
        {
            try
            {
                var blogDto = await _blogService.GetBlogsByStatusAsync(state);
                return CreateResponse(blogDto);
            }
            catch (Exception e)
            {
                return StatusCode(500, e.Message);
            }
        }

        [HttpGet("getByAuthor/{authorId:int}")]
        public ActionResult<List<BlogDto>> GetBlogsByAuthor(int authorId)
        {
            var result = _blogService.GetBlogsByAuthor(authorId);
            return CreateResponse(result);
        }

    }
}

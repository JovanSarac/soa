﻿using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.Public.Authoring;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System.Net.Http;

namespace Explorer.API.Controllers.Author.Authoring
{
    [Authorize(Policy = "authorPolicy")]
    [Route("api/tourManagement/tour")]
    public class TourController : BaseApiController
    {
        private readonly ITourService _tourService;
        private readonly HttpClient _httpClient;
        public TourController(ITourService tourService)
        {
            _httpClient = new HttpClient();
            _tourService = tourService;
        }

        [HttpGet]
        public ActionResult<PagedResult<TourDto>> GetAll([FromQuery] int page, [FromQuery] int pageSize)
        {
            var result = _tourService.GetPaged(page, pageSize);
            return CreateResponse(result);
        }

        //[HttpPost]
        //public ActionResult<TourDto> Create([FromBody] TourDto tour)
        //{
        //    var result = _tourService.Create(tour);
        //    return CreateResponse(result);
        //}
        [HttpPost]
        public async Task<ActionResult<string>> Create([FromBody] TourDto tour)
        {
            var result = await _tourService.CreateAsync(tour, _httpClient);
            return Ok(result);
        }

        [HttpPost("createTour")]
        public async Task<IActionResult> CreateAsync([FromBody] TourDto tour)
        {
            try
            {
                await _tourService.CreateAsync(tour);
                return Ok();
            }
            catch (Exception e)
            {
                return StatusCode(500, e.Message);
            }
        }



        //[HttpPut("{id:int}")]
        //public ActionResult<TourDto> Update([FromBody] TourDto tour)
        //{
        //    var result = _tourService.Update(tour);
        //    return CreateResponse(result);
        //}
        [HttpPut("{id:int}")]
        public async Task<string> Update([FromBody] TourDto tour)
        {
            var result = await _tourService.UpdateAsync(tour, _httpClient);
            return result;
        }


        [HttpPut("/updateTour/{id:int}")]
        public async Task<IActionResult> UpdateAsync([FromBody] TourDto tour)
        {
            try
            {
                await _tourService.UpdateAsync(tour);
                return Ok();
            }
            catch (Exception e)
            {
                return StatusCode(500, e.Message);
            }
        }

        [HttpDelete("{id:int}")]
        public ActionResult Delete(int id)
        {
            var result = _tourService.Delete(id);
            return CreateResponse(result);
        }

        [AllowAnonymous]
        [HttpGet("{id:int}")]
        public ActionResult<TourDto> Get(int id)
        {
            var result = _tourService.Get(id);
            return CreateResponse(result);
        }

        //[HttpPut("publish/{id:int}")]
        //public ActionResult<TourDto> Publish(int id, [FromBody] int authorId)
        //{
        //    var result = _tourService.Publish(id, authorId);
        //    return CreateResponse(result);
        //}

        [HttpPut("publish/{id:int}")]
        public async Task<ActionResult<string>> Publish(int id, [FromBody] int authorId)
        {
            var result = await _tourService.PublishAsync(id, authorId, _httpClient);
            return Ok(result);
        }

        //[HttpPut("archive/{id:int}")]
        //public ActionResult<TourDto> Archive(int id, [FromBody] int authorId)
        //{
        //    var result = _tourService.Archive(id, authorId);
        //    return CreateResponse(result);
        //}
        [HttpPut("archive/{id:int}")]
        public async Task<ActionResult<string>> Archive(int id, [FromBody] int authorId)
        {
            var result = await _tourService.ArchiveAsync(id, authorId, _httpClient);
            return Ok(result);
        }

        [HttpGet("author")]
        public ActionResult<PagedResult<TourDto>> GetAllByAuthorId([FromQuery] int authorId, [FromQuery] int page, [FromQuery] int pageSize)
        {
            var result = _tourService.GetPagedByAuthorId(authorId, page, pageSize);
            return CreateResponse(result);
        }
    }
}
